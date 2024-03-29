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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1ShareArticleRequest
 */
export interface V1ShareArticleRequest {
    /**
     * 記事のURL
     * @type {string}
     * @memberof V1ShareArticleRequest
     */
    url: string;
    /**
     * タイトル
     * @type {string}
     * @memberof V1ShareArticleRequest
     */
    title?: string;
    /**
     * description
     * @type {string}
     * @memberof V1ShareArticleRequest
     */
    description?: string;
    /**
     * サムネイルのURL
     * @type {string}
     * @memberof V1ShareArticleRequest
     */
    thumbnail?: string;
}

/**
 * Check if a given object implements the V1ShareArticleRequest interface.
 */
export function instanceOfV1ShareArticleRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "url" in value;

    return isInstance;
}

export function V1ShareArticleRequestFromJSON(json: any): V1ShareArticleRequest {
    return V1ShareArticleRequestFromJSONTyped(json, false);
}

export function V1ShareArticleRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ShareArticleRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'url': json['url'],
        'title': !exists(json, 'title') ? undefined : json['title'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'thumbnail': !exists(json, 'thumbnail') ? undefined : json['thumbnail'],
    };
}

export function V1ShareArticleRequestToJSON(value?: V1ShareArticleRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'url': value.url,
        'title': value.title,
        'description': value.description,
        'thumbnail': value.thumbnail,
    };
}

