import { Component, OnInit } from '@angular/core';
import { IdentificationNumberDataSource } from '../shared/identification-number.datasource';
import { IdentificationNumberService } from '../shared/identification-number.service';

@Component({
  selector: 'app-identification-number-list',
  templateUrl: './identification-number-list.component.html',
  styleUrls: ['./identification-number-list.component.css']
})
export class IdentificationNumberListComponent implements OnInit {

  displayedColumns: string[] = ['id', 'number', 'blocked'];

  constructor(private identificationNumberService:IdentificationNumberService,private identificationNumberDataSource: IdentificationNumberDataSource) {}

  ngOnInit(): void {
    this.identificationNumberDataSource.init();
  }

  get dataSource(){
    return this.identificationNumberDataSource;
  }
}
