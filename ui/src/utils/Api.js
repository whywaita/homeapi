import * as superagent from 'superagent';
import { Emit } from '../utils/EventBus';

async function RequestMiddleware(request) {
  if (request.url[0] !== '/' || request.url.startsWith('http')) {
    request.url = `/api/irkit/${request.url}`;
  }
  request.set('Content-Type', 'application/json');

  try {
    return await request;
  } catch (err) {
    if (err.status === 401 || err.status === 403) {
      Emit(err);
    }
    throw err;
  }
}

// eslint-disable-next-line
export class API {
  static getDeviceList(req = superagent) {
    return RequestMiddleware(req.get('list'))
      .then(res => res.body);
  }

  static changeDeviceOff(deviceID, req = superagent) {
    return RequestMiddleware(req.get(`${deviceID}/off`))
      .then(res => res.body);
  }

  static changeDeviceOn(deviceID, req = superagent) {
    return RequestMiddleware(req.get(`${deviceID}/on`))
      .then(res => res.body);
  }
}
