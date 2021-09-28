import { stringifyUrl, ParsedQuery } from 'query-string';
import { IFetchParams } from '../models/fetch/IFetchParams';
import { FetchMethod } from '../enums/FetchMethod';
import { IResponseError } from '../models/fetch/IResponseError';
import { env } from '../../env';

const getInitHeaders = (contentType = 'application/json', hasContent = true) => {
  const headers: HeadersInit = new Headers();

  if (hasContent) {
    headers.set('Content-Type', contentType);
  }
  return headers;
};

const getFetchUrl = (url: string, query: ParsedQuery) => stringifyUrl({ url, query }, { skipNull: true });

const getFetchOptions = (method: string, body?: IFetchParams) => ({
  method,
  headers: getInitHeaders(),
  body: body && JSON.stringify(body)
});

const throwIfResponseFailed = async (res: Response) => {
  if (!res.ok) {
    if (res.status === 401) {
      return;
    }
    let parsedException: IResponseError = {
      message: 'Something went wrong with request!',
      status: 500
    };
    try {
      parsedException = await res.json();
    } catch (err) {
      throw parsedException;
    }
  }
};

const makeRequest = (method: FetchMethod) => async <T>(url: string, params?: IFetchParams) => {
  const domainUrl = `${env.urls.server}${url}`;
  const [fetchUrl, body] = method === FetchMethod.GET
    ? [getFetchUrl(domainUrl, params as ParsedQuery), undefined]
    : [url, params];
  const fetchOptions = getFetchOptions(method, body);

  const res = await fetch(fetchUrl, fetchOptions);
  await throwIfResponseFailed(res);
  return res.json() as Promise<T>;
};

const api = {
  get: makeRequest(FetchMethod.GET),
  post: makeRequest(FetchMethod.POST),
  put: makeRequest(FetchMethod.PUT),
  delete: makeRequest(FetchMethod.DELETE)
};

export default api;