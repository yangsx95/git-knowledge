export enum ContentType {
  Text = "text/plain",
  Json = "application/json",
}

export class MessageBean {
  private _func: string;
  private _content_type: string;
  private _content: string;
  private _success: boolean;
  private _error_message: string;

  constructor(func: string, content_type: string, content: string) {
    this._func = func;
    this._content_type = content_type;
    this._content = content;
    this._success = true;
    this._error_message = "";
  }

  get func(): string {
    return this._func;
  }

  set func(value: string) {
    this._func = value;
  }

  get content_type(): string {
    return this._content_type;
  }

  set content_type(value: string) {
    this._content_type = value;
  }

  get content(): string {
    return this._content;
  }

  set content(value: string) {
    this._content = value;
  }

  get success(): boolean {
    return this._success;
  }

  set success(value: boolean) {
    this._success = value;
  }

  get error_message(): string {
    return this._error_message;
  }

  set error_message(value: string) {
    this._error_message = value;
  }

  toJSON(): any {
    return {
      func: this._func,
      content_type: this._content_type,
      content: this._content
    }
  }

  toJSONString(): string {
    return JSON.stringify(this.toJSON())
  }
}

/**
 * Websocket代理对象，用于包装消息、发送消息
 */
export class WebsocketProxy {

  private websocket: WebSocket;
  onopen: ((this: WebsocketProxy) => any) | null;
  onclose: ((this: WebsocketProxy) => any) | null;
  onerror: ((this: WebsocketProxy, errorMessage: string) => any) | null;
  onmessage: ((this: WebsocketProxy, message: any) => any) | null;

  constructor(url: string) {
    this.onopen = null;
    this.onclose = null;
    this.onmessage = null;
    this.onerror = null;
    this.websocket = new WebSocket(SERVER_ADDRESS_WS + url);
    this.websocket.onopen = () => {
      if (this.onopen) {
        this.onopen()
      }
    }
    this.websocket.onclose = () => {
      if (this.onclose) {
        this.onclose();
      }
    }
    this.websocket.onmessage = ev => {
      if (ev.data) {
        // 解析消息
        const m: MessageBean = JSON.parse(ev.data);
        let targetContent: any
        if (m.success) {
          switch (m.content_type) {
            case 'text/plain':
              targetContent = m.content.toString()
              break;
            case 'application/json':
              targetContent = JSON.parse(m.content.toString())
              break;
          }
          if (this.onmessage) this.onmessage(targetContent)
        } else {
          if (this.onerror) this.onerror(m.error_message)
        }
      }
    }
  }

  send(func: string, contentType: ContentType, content: any): void {
    switch (contentType) {
      case ContentType.Text:
        this.websocket.send(JSON.stringify(new MessageBean(func, contentType, content.toString())));
        break;
      case ContentType.Json: {
        this.websocket.send(JSON.stringify(new MessageBean(func, contentType, JSON.stringify(content))));
        break;
      }
    }
  }


}
