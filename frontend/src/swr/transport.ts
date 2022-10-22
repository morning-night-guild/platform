import { createConnectTransport } from '@bufbuild/connect-web';

// トランスポート作成
export const transport = createConnectTransport({
    baseUrl: import.meta.env.VITE_SERVICE_ENDPOINT,
});
