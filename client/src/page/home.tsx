import { useEffect, useRef, useState } from "react";
import Navbar from "../componets/navbar";
import gsap from "gsap";
import { useGSAP } from "@gsap/react";
import { Login } from "../componets/Login";


export const Home = () => {
	const animation = useRef(gsap.timeline({ paused: true }));
	const elementRef = useRef<HTMLDivElement>(null);
	gsap.registerPlugin(useGSAP);

	const [state, setState] = useState(false)
	useEffect(() => {
		animation.current = gsap.timeline({ paused: true })
			.fromTo(elementRef.current,
				{
					opacity: 0,
					x: 20,
					duration: 0.3,
				},
				{
					duration: 0.3,
					opacity: 1,
					x: 0,
					ease: "power1.in",
				}
			)
	}, []);

	const loginAtten = () => {
		if (state) {
			// Animate out and set visibi
			animation.current.reverse();
			setState(false)
		} else {
			// Set visibility to true and play the animation forward
			setState(true)
			animation.current.play();
		}
	}

	return (
		<>
			<Navbar />
			<h1 className="text-text-hed text-3xl flex  justify-center  ">store excess retive files</h1>
			<button className="text-text-hed" onClick={loginAtten}>sign in</button>
			<Login refElement={elementRef}/>


		</>
	);
};                                                                                                                                                                     
