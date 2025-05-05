import { PaginatedRespose } from 'src/app/shared/models/pagination.model';
import { Desenvolvedor } from './desenvolvedor.model';

export type CreateDesenvolvedorResponse = Desenvolvedor;

export type UpdateDesenvolvedorResponse = Desenvolvedor;

export type GetDesenvolvedoresResponse = PaginatedRespose<Desenvolvedor>;
