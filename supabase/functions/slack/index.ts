import { DOMParser, Element, NodeType } from "deno-dom";
import { serve } from "https://deno.land/std@0.131.0/http/server.ts";

type OGP = {
  title: string;
  description: string;
  thumbnail: string;
};

export type Env = {
  API_KEY: string;
  SLACK_VERIFICATION_TOKEN: string;
  CORE_SERVICE_URL: string;
};

export type SlackEventType = "url_verification" | "event_callback";

export type SlackEvent = {
  token: string;
  challenge: string;
  type: SlackEventType;
  event: Event;
};

export type EventType = "message";

export type EventSubType = "message_changed";

export type Event = {
  type: EventType;
  subtype: EventSubType;
  text: string;
  user: string;
  ts: number;
};

/**
 * '|'によって連結されたURL達から、最初のURLを取り出す
 *
 *  https://github.com/morning-night-guild/platform/issues/161 の暫定対応
 */
export const extractFirstUrlFromUrlsConcatByPipe = (url: string) => {
  if (!url?.includes("|")) {
    return url;
  }

  const urls = url.split("|");
  return urls[0];
};

const env: Env = {
  API_KEY: Deno.env.get("CORE_API_KEY") ?? "",
  CORE_SERVICE_URL: Deno.env.get("CORE_SERVICE_URL") ?? "",
  SLACK_VERIFICATION_TOKEN: Deno.env.get("SLACK_VERIFICATION_TOKEN") ?? "",
};

export const verify = (token1: string, token2: string) => {
  if (token1 !== token2) {
    throw new Error("token mismatch");
  }
};

export const callback = async (
  event: SlackEvent,
  endpoint: string,
  key: string,
) => {
  console.log(event.type);
  const pattern = /http(.*):\/\/([a-zA-Z0-9/\-\_\.]*)/;
  try {
    const u = event.event.text.match(pattern)?.find((s) => s);
    console.log(u);

    const url = extractFirstUrlFromUrlsConcatByPipe(u ?? "");

    const ogp = await createOGP(url);
    if (!ogp) {
      return;
    }
    const init = {
      body: JSON.stringify({
        url: url,
        title: ogp.title,
        description: ogp.description,
        thumbnail: ogp.thumbnail,
      }),
      method: "POST",
      headers: {
        "X-API-KEY": key,
        "Content-Type": "application/json",
      },
    };
    const res = await fetch(endpoint, init);

    console.debug(res);
  } catch (e) {
    console.warn(e);
    // 不要なリトライを防ぐため握りつぶす
    return;
  }
};

serve(async (request: { json: () => any }) => {
  const req = await request.json();

  const event: SlackEvent = JSON.parse(JSON.stringify(req));

  console.debug(JSON.stringify(event));

  verify(event.token, env.SLACK_VERIFICATION_TOKEN);

  if (event.type === "url_verification") {
    return new Response(JSON.stringify(event.challenge));
  }

  console.info(
    `received event type: ${event.type}, sub type: ${event.event.subtype}`,
  );

  if (
    event.type === "event_callback" && event.event.subtype === "message_changed"
  ) {
    return new Response(JSON.stringify(""));
  }

  if (event.type === "event_callback") {
    const url = env.CORE_SERVICE_URL + "/article.v1.ArticleService/Share";
    callback(event, url, env.API_KEY);
  }

  return new Response(
    JSON.stringify(""),
    { headers: { "Content-Type": "application/json" } },
  );
});

/**
 * OGPタグを取得して、そのcontentをJSON形式で返す.
 *
 * @param url URL
 */
const createOGP = async (url: string): Promise<OGP | undefined> => {
  try {
    const response = await fetch(url);

    const data = await response.text();
    const document = new DOMParser().parseFromString(data, "text/html")!;

    const meta = document?.querySelectorAll("head > meta")!;

    const ogp: { [key: string]: string } = {};

    // metaからOGPを抽出し、JSON形式に変換する
    Array.from(meta)
      .filter((node) => {
        return node.nodeType === NodeType.ELEMENT_NODE;
      })
      .map((node) => node as Element)
      .filter((element) => {
        return element.hasAttribute("property");
      })
      .forEach((element: Element) => {
        const property = element.getAttribute(
          "property",
        )?.trim().replace("og:", "");
        if (!property) {
          return;
        }
        const content = element.getAttribute(
          "content",
        );
        if (!content) {
          return content;
        }
        ogp[property] = content;
      });

    return {
      title: ogp["title"],
      description: ogp["description"],
      thumbnail: ogp["image"],
    };
  } catch (e) {
    console.warn(e);
  }
};
