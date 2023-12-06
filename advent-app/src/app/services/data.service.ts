import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { throwError } from 'rxjs';

const APIV1 = "v1";

export interface Data {
  answer1: string | number;
  answer2: string | number;
  metaData: any;
}

export interface AdventResponse {
  status: number;
  data: Data;
}

export interface AdventResponseError {
  status: number;
  error: string;
}

@Injectable({
  providedIn: 'root'
})
export class DataService {

  constructor(
    private http: HttpClient
  ) { }

  getDayAnswer(day: number) {
    return this.http.get<AdventResponse>(`http://localhost:8084/${APIV1}/advent/2022/${day}`)
      .pipe(
        catchError(error => {
          console.error(error);
          return throwError(() => error.error);
        })
      );
  }

  getDayAnswerForYear(year: number, day: number) {
    return this.http.get<AdventResponse>(`http://localhost:8084/${APIV1}/advent/${year}/${day}`)
      .pipe(
        catchError(error => {
          console.error(error);
          return throwError(() => error.error);
        })
      );
  }
}
