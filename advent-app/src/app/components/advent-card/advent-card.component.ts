import { Component, Input, OnInit } from '@angular/core';
import { DataService, Data, AdventResponse } from 'src/app/services/data.service';
import { finalize } from 'rxjs/operators';

@Component({
  selector: 'app-advent-card',
  templateUrl: './advent-card.component.html',
  styleUrls: ['./advent-card.component.scss']
})

export class AdventCardComponent {
  answers: Data | null = null;
  error: string = '';

  @Input() day: number = 0;
  @Input() title: string = '';
  @Input() loading: boolean = false;
  @Input() description: string = '';
  @Input() year: number = 0;

  constructor(private adventAPI: DataService) { }

  getDayAnswer() {
    this.loading = true;
    this.error = '';

    if (this.year !== null) {
      this.adventAPI.getDayAnswerForYear(this.year, this.day).pipe(
        finalize(() => this.loading = false)
      ).subscribe((answers: AdventResponse) => {
        this.answers = answers.data;
      }, error => {
        this.error = error;
      });

      return;
    }

    this.adventAPI.getDayAnswer(this.day).pipe(
      finalize(() => this.loading = false)
    ).subscribe((answers: AdventResponse) => {
      this.answers = answers.data;
    }, error => {
      this.error = error;
    });
  }
}
