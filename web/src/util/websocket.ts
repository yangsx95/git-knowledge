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
 * 消息观察者
 */
interface WebsocketObserver {
  /**
   * 是否需要通知
   */
  needResponse: () => boolean
  /**
   * 正确通知
   */
  response: (message: any) => void

  /**
   * 错误通知
   */
  errorResponse: (errorMessage: string) => void
}

/**
 * Websocket代理对象，用于包装消息、发送消息
 */
export class WebsocketSubject {

  private websocket: WebSocket;

  private observers: WebsocketObserver[] = [];

  constructor(url: string) {
    this.websocket = new WebSocket(SERVER_ADDRESS_WS + url);
    this.websocket.onopen = () => {
      console.log("websocket 建立了连接")
    }
    this.websocket.onclose = () => {
      console.log("websocket 关闭了连接")
    }
    this.websocket.onmessage = ev => {
      if (!ev.data) {
        return
      }
      const m: MessageBean = JSON.parse(ev.data);
      // 失败的消息
      if (!m.success) {
        this.observers.forEach(e => e.errorResponse(m.error_message))
        return;
      }
      // 成功的消息
      let targetContent: any
      switch (m.content_type) {
        case 'text/plain':
          targetContent = m.content.toString()
          break;
        case 'application/json':
          targetContent = JSON.parse(m.content.toString())
          break;
      }
      this.observers.forEach(e => e.response(targetContent))
    }
  }

  add(observer: WebsocketObserver) {
    this.observers.push(observer)
  }

  remove(observer: WebsocketObserver) {
    const index = this.observers.indexOf(observer)
    if (index < 0) {
      return
    }
    this.observers.splice(index, 1);
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
