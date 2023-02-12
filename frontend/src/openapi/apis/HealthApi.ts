/* tslint:disable */
/* eslint-disable */
/**
 * Morning Night Guild - App Gateway
 * This is the AppGateway API documentation.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: morning.night.guild@example.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';

/**
 * 
 */
export class HealthApi extends runtime.BaseAPI {

    /**
     * ヘルスチェック
     * apiヘルスチェック
     */
    async v1HealthAPIRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/health/api`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * ヘルスチェック
     * apiヘルスチェック
     */
    async v1HealthAPI(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1HealthAPIRaw(initOverrides);
    }

    /**
     * ヘルスチェック
     * coreヘルスチェック
     */
    async v1HealthCoreRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/health/core`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * ヘルスチェック
     * coreヘルスチェック
     */
    async v1HealthCore(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1HealthCoreRaw(initOverrides);
    }

}
