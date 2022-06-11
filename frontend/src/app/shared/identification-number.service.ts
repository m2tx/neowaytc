import { HttpClient, HttpHeaders, HttpRequest } from '@angular/common/http'
import { Injectable } from '@angular/core'
import { Observable } from 'rxjs'
import { environment } from 'src/environments/environment'
import { IdentificationNumber } from './identification-number.model';
import { Page } from './page.model';

const HTTP_OPTIONS = {
  headers: new HttpHeaders(
    {
      'Content-Type': 'application/json',
    }
  )
};

@Injectable({
  providedIn: 'root'
})
export class IdentificationNumberService {

  constructor(
    private http: HttpClient
  ) {}

  create(identificationNumber: IdentificationNumber): Observable<IdentificationNumber> {
    return this.http.post<IdentificationNumber>(`${environment.api}/`, document, HTTP_OPTIONS)
  }

  findById(id: string): Observable<IdentificationNumber> {
    return this.http.get<IdentificationNumber>(`${environment.api}/${id}`);
  }

  queryBy(params: IdentificationNumber, column: string = 'number', sort: string = 'asc', page: number = 0, size: number = 5): Observable<Page<IdentificationNumber>> {
    return this.http.post<Page<IdentificationNumber>>(`${environment.api}/?sort=${column},${sort}&page=${page}&size=${size}`, params);
  }

}