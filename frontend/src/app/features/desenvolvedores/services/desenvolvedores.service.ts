import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.development';
import { CreateDesenvolvedorRequest, UpdateDesenvolvedorRequest } from '../models/request.model';
import { CreateDesenvolvedorResponse, GetDesenvolvedoresResponse, UpdateDesenvolvedorResponse } from '../models/response.model';

@Injectable({
  providedIn: 'root',
})
export class DesenvolvedoresService {
  private api = environment.backend;

  constructor(private http: HttpClient) {}

  getDesenvolvedores(page: number = 1, search: string = '', order: string = '', mode: string = ''): Observable<GetDesenvolvedoresResponse> {
    let params = new HttpParams();

    if (page > 1) params = params.set('page', page.toString());
    if (order != '') params = params.set('order', order);
    if (mode != '') params = params.set('mode', mode);
    if (search != '') params = params.set('search', search);

    return this.http.get<GetDesenvolvedoresResponse>(this.api(`/desenvolvedores`), { params });
  }

  createDesenvolvedor(desenvolvedor: CreateDesenvolvedorRequest): Observable<CreateDesenvolvedorResponse> {
    return this.http.post<CreateDesenvolvedorResponse>(this.api(`/desenvolvedores`), desenvolvedor);
  }

  updateDesenvolvedor(id: string, desenvolvedor: UpdateDesenvolvedorRequest): Observable<UpdateDesenvolvedorResponse> {
    return this.http.put<UpdateDesenvolvedorResponse>(this.api(`/desenvolvedores/${id}`), desenvolvedor);
  }

  deleteDesenvolvedor(id: string): Observable<void> {
    return this.http.delete<void>(this.api(`/desenvolvedores/${id}`));
  }
}
