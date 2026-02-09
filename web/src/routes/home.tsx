import { Header } from '@/components/home/header'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/home')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <main className="min-h-screen text-white bg-zinc-500">

        <Header/>

      <div className=" w-full grid grid-cols-2 md:grid-cols-[1fr_3fr_1fr] gap-4">
        <aside className="bg-blue-900 ">ASIDE</aside>

        <section className="bg-green-900">SECTION</section>

        <aside className="bg-yellow-900">ASIDE</aside>
      </div>
    </main>
  )
}
