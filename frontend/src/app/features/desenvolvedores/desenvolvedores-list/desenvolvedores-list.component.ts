import { Component } from '@angular/core';
import { Desenvolvedor, DesenvolvedorForm } from '../models/desenvolvedor.model';
import { Pagination } from 'src/app/shared/models/pagination.model';
import { ModalService } from 'src/app/shared/components/modal/modal.service';
import { ToastService } from 'src/app/shared/components/toast/toast.service';
import { DesenvolvedoresService } from '../services/desenvolvedores.service';
import { GetDesenvolvedoresResponse } from '../models/response.model';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-desenvolvedores-list',
  templateUrl: './desenvolvedores-list.component.html',
  styleUrls: ['./desenvolvedores-list.component.scss'],
})
export class DesenvolvedoresListComponent {
  desenvolvedores: Desenvolvedor[] = [];
  search: string = '';
  pagination: Pagination = {
    total: 1,
    per_page: 10,
    current_page: 1,
    last_page: 1,
  };

  showForm = false;
  editingDesenvolvedor: DesenvolvedorForm | null = null;

  orderField = '';
  orderDirection: 'ASC' | 'DESC' = 'ASC';

  columns = [
    { label: 'Nome', field: 'nome', hidden: false },
    { label: 'Sexo', field: 'sexo', hidden: false },
    { label: 'Data de Nascimento', field: 'data_nascimento', hidden: false },
    { label: 'Hobby', field: 'hobby', hidden: false },
    { label: 'Nível', field: 'nivel.nivel', hidden: false },
  ];

  constructor(
    private desenvolvedoresService: DesenvolvedoresService,
    private toastService: ToastService,
    private modalService: ModalService,
  ) {}

  ngOnInit(): void {
    this.fetchDesenvolvedores();
  }

  onSearch(search: string) {
    this.search = search;
    this.fetchDesenvolvedores(1, search);
  }

  onOrderChange(order: { field: string; direction: 'ASC' | 'DESC' }) {
    this.orderField = order.field;
    this.orderDirection = order.direction;
    this.fetchDesenvolvedores(this.pagination.current_page, this.search, order.field, order.direction);
  }

  onPageChange(page: number) {
    this.fetchDesenvolvedores(page, this.search, this.orderField, this.orderDirection);
  }

  fetchDesenvolvedores(page: number = 1, search: string = '', order: string = '', mode: string = '') {
    this.desenvolvedoresService.getDesenvolvedores(page, search, order, mode).subscribe({
      next: (response: GetDesenvolvedoresResponse) => {
        this.desenvolvedores = response.data;
        this.pagination = response.meta;
      },
      error: () => {
        this.toastService.show('Não há itens para listar', true);
      },
    });
  }

  onCreate() {
    const desenvolvedor: DesenvolvedorForm = { id: '', nome: '', sexo: '', data_nascimento: '', hobby: '', nivel_id: '' };
    this.openForm(desenvolvedor);
  }

  openForm(desenvolvedor: DesenvolvedorForm) {
    this.editingDesenvolvedor = desenvolvedor;
    this.showForm = true;
  }

  closeForm(desenvolvedor: DesenvolvedorForm | null) {
    this.editingDesenvolvedor = null;
    this.showForm = false;

    if (!desenvolvedor) return;

    if (desenvolvedor.id) {
      this.updateDesenvolvedor(desenvolvedor);
    } else {
      this.createDesenvolvedor(desenvolvedor);
    }
  }

  private createDesenvolvedor(desenvolvedor: DesenvolvedorForm) {
    this.desenvolvedoresService.createDesenvolvedor(desenvolvedor).subscribe({
      next: () => {
        this.toastService.show('Desenvolvedor criado com sucesso');
        this.fetchDesenvolvedores();
      },
      error: () => {
        this.toastService.show('Erro ao criar desenvolvedor', true);
      },
    });
  }

  private updateDesenvolvedor(desenvolvedor: DesenvolvedorForm) {
    this.desenvolvedoresService.updateDesenvolvedor(desenvolvedor.id, desenvolvedor).subscribe({
      next: () => {
        this.toastService.show('Desenvolvedor atualizado com sucesso');
        this.fetchDesenvolvedores();
      },
      error: () => {
        this.toastService.show('Erro ao atualizar desenvolvedor', true);
      },
    });
  }

  async onDelete(desenvolvedor: Desenvolvedor) {
    const confirmed = await this.modalService.confirm('Deseja realmente excluir este desenvolvedor?');

    if (!confirmed) return;

    this.desenvolvedoresService.deleteDesenvolvedor(desenvolvedor.id).subscribe({
      next: () => {
        this.toastService.show('Desenvolvedor excluído com sucesso');
        this.fetchDesenvolvedores();
      },
      error: () => {
        this.toastService.show('Erro ao excluir o desenvolvedor', true);
      },
    });
  }
}
