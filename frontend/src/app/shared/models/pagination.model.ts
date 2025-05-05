export interface PaginationMeta {
  total: number;
  per_page: number;
  current_page: number;
  last_page: number;
}

export interface PaginatedRespose<T> {
  data: T[];
  meta: PaginationMeta;
}

export type Pagination = PaginationMeta;
