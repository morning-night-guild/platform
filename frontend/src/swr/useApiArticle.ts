import useSWR from 'swr';

import { createPromiseClient } from '@bufbuild/connect-web';
import { ArticleService } from '../api/connect/proto/article/v1/article_connectweb';
import type { Article } from '../api/connect/proto/article/v1/article_pb';
import { transport } from './transport';

// 1画面に表示する記事の最大数
const maxPageSize = 100;

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
    const request = {
        maxPageSize,
        pageToken: articlesState.currentIndex,
    };

    const fetcher = async () => client.list(request);

    const { data } = useSWR('/api/v1/articles', fetcher);
    articlesState.data = data?.articles ?? [];
    articlesState.currentIndex = data?.nextPageToken;

    return data;
};
