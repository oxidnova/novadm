import { requestClient } from '#/api/request';

export namespace ConsultationApi {
  export interface FetchParams {
    page: number;
    pageSize: number;
    startTime?: number;
    endTime?: number;
    id?: string;
    status?: string;
  }

  export interface Consultation {
    id: string;
    prompt: string;
    content: string;
    status: number;
    createdAt: string;
    updatedAt: string;
  }
}

async function searchConsultationsApi(params: ConsultationApi.FetchParams) {
  return requestClient.get('/consultation', { params });
}

export { searchConsultationsApi };
