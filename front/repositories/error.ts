export class APIStatusError implements Error {
  name: string = "status code error";
  message: string = "failed api request";
  stack?: string;

  toString() {
    return `${this.name}:${this.message}`;
  }
}
