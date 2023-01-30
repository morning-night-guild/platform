import useSWR, { useSWRConfig } from 'swr';

import { createPromiseClient } from '@bufbuild/connect-web';
import { ArticleService } from '../api/connect/proto/article/v1/article_connectweb';
import type { Article } from '../api/connect/proto/article/v1/article_pb';
import { transport } from './transport';

// 1画面に表示する記事の数
const articlesPerPage = 20;

// クライアント作成
const client = createPromiseClient(ArticleService, transport);

type ArticlesState = {
    data: Article[];
    error?: Error;
    currentIndex?: string;
};

const articlesState: ArticlesState = {
    data: [],
};

export const useListArticles = () => {
    const key = `/api/v1/articles`;

    const request = {
        articlesPerPage,
        pageToken: articlesState.currentIndex,
    };

    const fetcher = async () => client.list(request);

    const { data } = useSWR(key, fetcher);

    const fetchedArticles = data?.articles ?? [];
    const existIds = articlesState.data.map(d => d.id);
    const additionalArticles = fetchedArticles.filter(d => existIds.indexOf(d.id) < 0);
    articlesState.data.push(...additionalArticles);

    // NextPageTokenが空の場合、もうこれ以上データがないのでcurrentIndexを更新しない
    if (!(data?.nextPageToken === '')) {
        articlesState.currentIndex = data?.nextPageToken;
    }

    const { mutate } = useSWRConfig();

    return {
        data: articlesState.data,
        async mutate() {
            await mutate(key);
        },
    };
};
