import { Component } from '@angular/core';
interface Day {
  dayNumber: number;
  title: string;
  description: string;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'advent-app';

  days: Day[] = [{
    dayNumber: 1,
    title: 'Calories Count',
    description: 'Sum the total calories'
  },{
    dayNumber: 2,
    title: 'Rock Paper Scissors',
    description: 'Use proper strategy to win!'
  },{
    dayNumber: 3,
    title: 'Rucksack Reorganization',
    description: 'Sum the priorities'
  },{
    dayNumber: 4,
    title: 'Camp Cleanup',
    description: 'What overlaps!!'
  },{
    dayNumber: 5,
    title: 'Supply Stacks',
    description: 'What ends up on the top of the stacks!!'
  },{
    dayNumber: 6,
    title: 'Tuning Trouble',
    description: 'Find the start of the packet marker!!'
  },{
    dayNumber: 7,
    title: 'No Space Left On Device',
    description: 'Find the correct directory to delete!'
  },{
    dayNumber: 8,
    title: 'Treetop Tree House',
    description: 'Find the best location for the tree house!'
  },{
    dayNumber: 9,
    title: 'Rope Bridge',
    description: 'Snake'
  },{
    dayNumber: 10,
    title: 'Cathode-Ray Tube',
    description: 'CRT and CPU cycles!'
  }];
}
