export interface Nivel {
  id: string;
  nivel: string;
  qtd_devs: number;
}

export type NivelForm = Omit<Nivel, 'qtd_devs'>;

export type NivelSelectOptions = Omit<Nivel, 'qtd_devs'>;
