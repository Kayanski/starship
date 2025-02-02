import { relative } from 'path';
import strip from 'strip-ansi'

import { StarshipClient } from '../src'
import { config } from './config';

interface TestClientCtx {
  commands: string[];
  logs: string[];
  code: number;
}

export const createClient = () => {
  const ctx: TestClientCtx = {
    commands: [],
    logs: [],
    code: -1
  };

  const client = new StarshipClient({
    helmName: 'osmojs',
    helmFile: relative(process.cwd(), config.configPath),
  });

  const handler = {
    get(target: any, prop: string, _receiver: any) {
      const originalMethod = target[prop];
      if (typeof originalMethod === 'function') {
        return function(...args: any[]) {
          const argsString = args.map(arg => strip(JSON.stringify(arg))).join(', ');
          ctx.logs.push(`Call: ${prop}(${argsString})`);
          // if you want to see nested Call, replace target with this
          // I double checked both this and target, it does not call the exec in the methods when used internally
          return originalMethod.apply(target, args);
        };
      }
      return originalMethod;
    }
  };

  const proxiedClient: StarshipClient = new Proxy(client, handler);

  // Overriding the exit method
  // @ts-ignore
  proxiedClient.exit = (code: number) => {
    ctx.code = code;
  };
  
  // @ts-ignore
  proxiedClient.ensureFileExists = (_filename: string) => {
    
  };
  
  // Overriding the exec method
  // @ts-ignore
  proxiedClient.exec = (cmd: string[]) => {
    // @ts-ignore
    client.checkDependencies();
    ctx.commands.push(cmd.join(' '));
    ctx.logs.push(cmd.join(' '));
    return '';
  };

  // Overriding the log method
  // @ts-ignore
  proxiedClient.log = (cmd: string) => {
    const str = strip(cmd);
    if (/\n/.test(str)) {
      ctx.logs.push('Log⬇\n' + str + '\nEndLog⬆');
    } else {
      ctx.logs.push('Log: ' + str);
    }
  };

  return {
    client: proxiedClient,
    ctx
  }
}


export const expectClient = (ctx: TestClientCtx, code: number) => {
  expect(ctx.logs.join('\n')).toMatchSnapshot();
  expect(ctx.commands.join('\n')).toMatchSnapshot();
  expect(ctx.code).toBe(code);
}
