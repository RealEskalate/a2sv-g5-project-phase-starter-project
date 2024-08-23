export interface Service{
      id: string,
      name: string,
      details: string,
      numberOfUsers: number,
      status: string,
      type: string,
      icon: string
}

export interface ServicePostRequest{
    name: string,
    details: string,
    numberOfUsers: number,
    status: string,
    type: string,
    icon: string
  }

export interface ServiceResponce{
    success: true,
    message: string,
    data:{
      content:Service[]
    }
    totalPages:number
  }