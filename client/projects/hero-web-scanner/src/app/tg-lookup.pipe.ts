import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'tgLookup'
})
export class TgLookupPipe implements PipeTransform {

  transform(value: unknown, ...args: unknown[]): unknown {
    return null;
  }

}
