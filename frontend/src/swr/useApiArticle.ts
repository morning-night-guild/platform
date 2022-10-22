import useSWR from 'swr';

import { ArticleService } from '../api/article/v1/article_connectweb';
import { Article } from '../api/article/v1/article_pb';
import { createPromiseClient } from '@bufbuild/connect-web';
import { transport } from './transport';

// 1画面に表示する記事の最大数
const MAX_PAGE_SIZE = 100;

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
    const req = {
        maxPageSize: MAX_PAGE_SIZE,
        pageToken: articlesState.currentIndex,
    };

    const fetcher = () => client.list(req);

    const { data, error } = useSWR('/api/v1/articles', fetcher);
    articlesState.data = data?.articles ?? [];
    articlesState.currentIndex = data?.nextPageToken;
    articlesState.error = error;

    return {
        data,
        error,
    };
};
