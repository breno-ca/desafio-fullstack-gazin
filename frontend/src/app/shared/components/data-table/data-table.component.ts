import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Pagination } from '../../models/pagination.model';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-data-table',
  templateUrl: './data-table.component.html',
  styleUrls: ['./data-table.component.scss'],
  standalone: true,
  imports: [CommonModule],
})
export class DataTableComponent {
  @Input() columns: { label: string; field: string; hidden?: boolean }[] = [];
  @Input() data: any[] = [];
  @Input() search: string = '';
  @Input() pagination: Pagination = {
    total: 0,
    per_page: 10,
    current_page: 1,
    last_page: 1,
  };

  orderField: string = '';
  orderDirection: 'ASC' | 'DESC' = 'ASC';

  @Output() create = new EventEmitter<void>();
  @Output() edit = new EventEmitter<any>();
  @Output() delete = new EventEmitter<any>();
  @Output() searchChange = new EventEmitter<string>();
  @Output() pageChange = new EventEmitter<number>();
  @Output() orderChange = new EventEmitter<{ field: string; direction: 'ASC' | 'DESC' }>();

  order(field: string) {
    if (this.orderField === field) {
      this.orderDirection = this.orderDirection === 'ASC' ? 'DESC' : 'ASC';
    } else {
      this.orderField = field;
      this.orderDirection = 'ASC';
    }

    this.orderChange.emit({ field: this.orderField, direction: this.orderDirection });
  }

  previous() {
    if (this.pagination.current_page > 1) {
      this.pageChange.emit(this.pagination.current_page - 1);
    }
  }

  next() {
    if (this.pagination.current_page < this.pagination.last_page) {
      this.pageChange.emit(this.pagination.current_page + 1);
    }
  }

  getFieldValue(row: any, field: string): any {
    return field.includes('.') ? field.split('.').reduce((obj, key) => obj?.[key], row) : row?.[field];
  }
}
