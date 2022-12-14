// @generated by protoc-gen-connect-web v0.2.1 with parameter "target=ts"
// @generated from file proto/health/v1/health.proto (package health.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {CheckRequest, CheckResponse} from "./health_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * ヘルスサービス
 *
 * @generated from service health.v1.HealthService
 */
export const HealthService = {
  typeName: "health.v1.HealthService",
  methods: {
    /**
     * チェック
     * Need X-Api-Key Header
     *
     * @generated from rpc health.v1.HealthService.Check
     */
    check: {
      name: "Check",
      I: CheckRequest,
      O: CheckResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

