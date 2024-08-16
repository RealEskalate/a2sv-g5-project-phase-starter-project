import { useGetAllTransactionQuery } from "@/lib/service/TransactionService";
import { Item } from "./LastTransaction";

export async function getLastTransaction(): Promise<Item[]> {
  const access =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJzYW1pdGVzdCIsImlhdCI6MTcyMzgwNzQzNiwiZXhwIjoxNzIzODkzODM2fQ.dTGZWeZVDP1btw1nf_hW84Zr5CPjW32hnj-vXlsWCUQz4MlU1EuTvHhSp3-xfmUZ";
  const { data, isError, isLoading, isSuccess } =
    await useGetAllTransactionQuery(access);

  if (isLoading) {
    console.log("fetching");
  }

  if (isError) {
    console.log("error");
  }

  let retrunData: Item[] = [];

  if (isSuccess) {
    retrunData = data.data;
  }
  return retrunData;
}

export const items: Promise<Item[]> = getLastTransaction();
