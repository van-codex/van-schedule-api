import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';

import { CronService } from './common/cron.service';
import { DbService } from './common/db.service';

import { IndexController } from './app/index.controller';
import { ScriptTaskController } from './app/script-task.controller';

import { ScriptTaskEntity } from './database/script-task.entity';
import { ApiTaskEntity } from './database/api-task.entity';
import { QueueTaskEntity } from './database/queue-task.entity';

@Module({
  imports: [
    TypeOrmModule.forRoot(),
    TypeOrmModule.forFeature([
      ScriptTaskEntity,
      ApiTaskEntity,
      QueueTaskEntity,
    ]),
  ],
  controllers: [
    IndexController,
    ScriptTaskController,
  ],
  providers: [
    DbService,
    CronService,
  ],
})
export class AppModule {
}
