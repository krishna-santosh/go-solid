import { Route, Routes } from '@solidjs/router'
import Navbar from './components/Navbar'
import Home from './pages/Home'
import Info from './pages/Info'

function App() {
	return (
		<>
			<Navbar />
			<div class="flex h-screen w-screen items-center justify-center bg-base-100">
				<Routes>
					<Route path={'/'} component={Home} />
					<Route path={'/info'} component={Info} />
				</Routes>
			</div>
		</>
	)
}

export default App
