import { requestClient } from '#/api/request';

export namespace ConsultationApi {
  export interface FetchParams {
    page: number;
    pageSize: number;
    startTime: number;
    endTime: number;
    status: string;
  }

  export interface Consultation {
    status: string;
  }
}

async function searchConsultationsApi(params: ConsultationApi.FetchParams) {
  return requestClient.get('/ses', { params });
}

export { searchConsultationsApi };
