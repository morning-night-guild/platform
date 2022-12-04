// @generated by protoc-gen-connect-web v0.2.1 with parameter "target=ts"
// @generated from file proto/article/v1/article.proto (package proto.article.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {ListRequest, ListResponse, ShareRequest, ShareResponse} from "./article_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * 記事サービス
 *
 * @generated from service proto.article.v1.ArticleService
 */
export const ArticleService = {
  typeName: "proto.article.v1.ArticleService",
  methods: {
    /**
     * 共有
     * Need X-Api-Key Header
     *
     * @generated from rpc proto.article.v1.ArticleService.Share
     */
    share: {
      name: "Share",
      I: ShareRequest,
      O: ShareResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 一覧
     *
     * @generated from rpc proto.article.v1.ArticleService.List
     */
    list: {
      name: "List",
      I: ListRequest,
      O: ListResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

