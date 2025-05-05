export const environment = {
  host: '',
  pathPrefix: '/api',

  backend(path: string): string {
    return `${environment.host}${environment.pathPrefix}${path}`;
  },
};
