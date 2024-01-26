import TaskInput from "@/components/TaskHeader";
import TodoTabs from "@/components/TodoTabs";

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

async function getData() {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks`);

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to fetch data");
  }

  return res.json();
}
