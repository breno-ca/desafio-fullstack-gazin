import { Component, OnInit } from '@angular/core';
import { Nivel, NivelForm } from '../models/nivel.model';
import { NiveisService } from '../services/niveis.service';
import { Pagination } from 'src/app/shared/models/pagination.model';
import { GetNiveisResponse } from '../models/response.model';
import { ToastService } from 'src/app/shared/components/toast/toast.service';
import { ModalService } from 'src/app/shared/components/modal/modal.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-niveis-list',
  templateUrl: './niveis-list.component.html',
})
export class NiveisListComponent implements OnInit {
  niveis: Nivel[] = [];
  search: string = '';
  pagination: Pagination = {
    total: 1,
    per_page: 10,
    current_page: 1,
    last_page: 1,
  };

  showForm = false;
  editingNivel: NivelForm | null = null;

  orderField = '';
  orderDirection: 'ASC' | 'DESC' = 'ASC';

  columns = [
    { label: 'ID', field: 'id', hidden: true },
    { label: 'Nivel', field: 'nivel', hidden: false },
    { label: 'Qtd. Devs', field: 'qtd_devs', hidden: false },
  ];

  constructor(
    private niveisService: NiveisService,
    private toastService: ToastService,
    private modalService: ModalService,
  ) {}

  ngOnInit(): void {
    this.fetchNiveis();
  }

  onSearch(search: string) {
    this.search = search;
    this.fetchNiveis(1, search);
  }

  onOrderChange(order: { field: string; direction: 'ASC' | 'DESC' }) {
    this.orderField = order.field;
    this.orderDirection = order.direction;
    this.fetchNiveis(this.pagination.current_page, this.search, order.field, order.direction);
  }

  onPageChange(page: number) {
    this.fetchNiveis(page, this.search, this.orderField, this.orderDirection);
  }

  fetchNiveis(page: number = 1, search: string = '', order: string = '', mode: string = '') {
    this.niveisService.getNiveis(page, search, order, mode).subscribe({
      next: (response: GetNiveisResponse) => {
        this.niveis = response.data;
        this.pagination = response.meta;
      },
      error: () => {
        this.toastService.show('Não há itens para listar', true);
      },
    });
  }

  onCreate() {
    const nivel: NivelForm = { id: '', nivel: '' };

    this.openForm(nivel);
  }

  async onDelete(nivel: Nivel) {
    const confirmed = await this.modalService.confirm('Deseja realmente excluir este nível?');

    if (!confirmed) return;

    this.niveisService.deleteNivel(nivel.id).subscribe({
      next: () => {
        this.toastService.show('Nível excluído com sucesso');
        this.fetchNiveis();
      },
      error: (error: HttpErrorResponse) => {
        if (error.status === 400) {
          this.toastService.show('Não é possível excluir um nivel com desenvolvedores associados', true);
          return;
        }

        this.toastService.show('Erro ao excluir o nível', true);
        console.error(error.message);
      },
    });
  }

  openForm(nivel: NivelForm) {
    this.editingNivel = nivel;
    this.showForm = true;
  }

  closeForm(nivel: NivelForm | null) {
    this.editingNivel = null;
    this.showForm = false;

    if (!nivel) return;

    if (nivel.id) {
      this.updateNivel(nivel);
    } else {
      this.createNivel(nivel);
    }
  }

  private createNivel(nivel: NivelForm) {
    this.niveisService.createNivel(nivel).subscribe({
      next: () => {
        this.toastService.show('Nível criado com sucesso');
        this.fetchNiveis();
      },
      error: () => {
        this.toastService.show('Erro ao criar o nível', true);
      },
    });
  }

  private updateNivel(nivel: NivelForm) {
    this.niveisService.updateNivel(nivel.id!, nivel).subscribe({
      next: () => {
        this.toastService.show('Nível atualizado com sucesso');
        this.fetchNiveis();
      },
      error: () => {
        this.toastService.show('Erro ao atualizar o nível', true);
      },
    });
  }
}
