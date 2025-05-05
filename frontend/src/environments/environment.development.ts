export const environment = {
  host: 'http://localhost:8080',
  pathPrefix: '/api',

  backend(path: string): string {
    return `${environment.host}${environment.pathPrefix}${path}`;
  },
};
