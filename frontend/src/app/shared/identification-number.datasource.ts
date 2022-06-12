import { CollectionViewer, DataSource } from '@angular/cdk/collections';
import { Injectable } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { PageEvent } from '@angular/material/paginator';
import { Sort } from '@angular/material/sort';
import { BehaviorSubject, Observable, of } from 'rxjs';
import { catchError, finalize } from 'rxjs/operators';
import { IdentificationNumber } from './identification-number.model';
import { IdentificationNumberService } from './identification-number.service';
import { Page } from './page.model';

@Injectable({
    providedIn: 'root'
})
export class IdentificationNumberDataSource implements DataSource<IdentificationNumber> {

  public form: FormGroup = new FormGroup({
    number: new FormControl(),
    blocked: new FormControl()
  });
  
  sort: Sort = {active: 'number', direction: 'desc'};
  public length: number = 0;
  public size: number = 10;
  
  private listSubject = new BehaviorSubject<IdentificationNumber[]>([]);
  private loadingSubject = new BehaviorSubject<boolean>(false);
  public loading$ = this.loadingSubject.asObservable();

  constructor(private identificationNumberService: IdentificationNumberService) {}

  init() {
    this.load(this.form.value, this.sort.active, this.sort.direction, 0, this.size);
  }

  load(params: IdentificationNumber, column: string, sort: string, page: number, size: number) {
    this.loadingSubject.next(true);
    this.identificationNumberService.queryBy(params, column, sort, page, size)
      .pipe(
        catchError(() => of([])),
        finalize(() => this.loadingSubject.next(false))
      )
      .subscribe(data => {        
        let page = data as Page<IdentificationNumber>;
        this.length = page.totalElements;
        this.size = page.size;
        this.listSubject.next(page.content)
      });
  }

  setPage(event: PageEvent) {
    this.load(this.form.value, this.sort.active, this.sort.direction, event.pageIndex, event.pageSize);
  }

  setSort(event: Sort) {
    this.sort = event;
    this.load(this.form.value, event.active, event.direction, 0, this.size);
  }

  connect(collectionViewer: CollectionViewer): Observable<IdentificationNumber[]> {
    return this.listSubject.asObservable();
  }

  disconnect(collectionViewer: CollectionViewer): void {
    this.listSubject.complete();
    this.loadingSubject.complete();
  }
}