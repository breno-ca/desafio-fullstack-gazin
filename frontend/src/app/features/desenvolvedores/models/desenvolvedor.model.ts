import { Nivel } from '../../niveis/models/nivel.model';

export interface Desenvolvedor {
  id: string;
  nivel_id: string;
  nome: string;
  sexo: string;
  data_nascimento: string;
  hobby: string;
  nivel: Omit<Nivel, 'qtd_devs'>;
}

export type DesenvolvedorForm = Omit<Desenvolvedor, 'nivel'>;
