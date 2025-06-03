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

async function genConsultationsApi(prompt: string) {
  return requestClient.post('/consultation/gen', { params: { prompt } });
}

async function updateConsultationsApi(data: ConsultationApi.Consultation) {
  return requestClient.post('/consultation/gen', data);
}

async function deleteConsultationsApi(id: string) {
  return requestClient.delete(`/consultation/${id}`);
}

export {
  deleteConsultationsApi,
  genConsultationsApi,
  searchConsultationsApi,
  updateConsultationsApi,
};
