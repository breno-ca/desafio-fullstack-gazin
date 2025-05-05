import { PaginatedRespose } from 'src/app/shared/models/pagination.model';
import { Nivel } from './nivel.model';

export type CreateNivelResponse = Nivel;

export type UpdateNivelResponse = Nivel;

export type GetNiveisResponse = PaginatedRespose<Nivel>;

export type GetSelectOptions = Omit<Nivel, 'qtd_devs'>[];
