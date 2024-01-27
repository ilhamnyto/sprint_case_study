"use client";

import TaskInput from "@/components/TaskHeader";
import TodoTabs from "@/components/TodoTabs";

export default function ErrorBoundary({ error }) {
  return (
    <>
      <div className="bg-red-400 px-3 py-2 rounded text-white font-semibold">
        {error.message} {": can't connect to API"}
      </div>
      <main className="grid place-items-center h-screen bg-slate-100">
        <div className="space-y-2 px-3 w-full max-w-xl">
          <div className="bg-white w-full min-h-[15vh] p-6 shadow rounded">
            <TaskInput />
          </div>
          <div className="">
            <TodoTabs datas={[]} />
          </div>
        </div>
      </main>
    </>
  );
}
