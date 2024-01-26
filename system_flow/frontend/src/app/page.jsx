import TaskInput from "@/components/TaskHeader";
import TodoTabs from "@/components/TodoTabs";
import { getData } from "@/lib/utils";
import { useContext } from "react";

export const dynamic = "force-dynamic";

export default async function Home() {
  const datas = await getData();

  return (
    <main className="grid place-items-center h-screen bg-slate-100">
      <div className="space-y-2 px-3 w-full max-w-xl">
        <div className="bg-white w-full min-h-[15vh] p-6 shadow rounded">
          <TaskInput />
        </div>
        <div className="">
          <TodoTabs datas={datas} />
        </div>
      </div>
    </main>
  );
}
