import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.development';
import { CreateNivelRequest, UpdateNivelRequest } from '../models/request.model';
import { CreateNivelResponse, GetNiveisResponse, GetSelectOptions, UpdateNivelResponse } from '../models/response.model';

@Injectable({
  providedIn: 'root',
})
export class NiveisService {
  private api = environment.backend;

  constructor(private http: HttpClient) {}

  getNiveis(page: number = 1, search: string = '', order: string = '', mode: string = ''): Observable<GetNiveisResponse> {
    let params = new HttpParams();

    if (page > 1) params = params.set('page', page.toString());
    if (order != '') params = params.set('order', order);
    if (mode != '') params = params.set('mode', mode);
    if (search != '') params = params.set('search', search);

    return this.http.get<GetNiveisResponse>(this.api(`/niveis`), { params });
  }

  getSelectOptions(): Observable<GetSelectOptions> {
    return this.http.get<GetSelectOptions>(this.api(`/niveis/select-options`));
  }

  createNivel(nivel: CreateNivelRequest): Observable<CreateNivelResponse> {
    return this.http.post<CreateNivelResponse>(this.api(`/niveis`), nivel);
  }

  updateNivel(id: string, nivel: UpdateNivelRequest): Observable<UpdateNivelResponse> {
    return this.http.put<UpdateNivelResponse>(this.api(`/niveis/${id}`), nivel);
  }

  deleteNivel(id: string): Observable<void> {
    return this.http.delete<void>(this.api(`/niveis/${id}`));
  }
}
