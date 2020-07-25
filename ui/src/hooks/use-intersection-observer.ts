import React from "react";


interface useIntersectionObserverProps {
    target: any;
    onIntersect: any;
    threshold?: number
    rootMargin?: string
}


const useIntersectionObserver = ({
    target,
    onIntersect,
    threshold = 0.1,
    rootMargin = "0px"
}: useIntersectionObserverProps) => {
    React.useEffect(() => {
        const observer = new IntersectionObserver(onIntersect, {
            rootMargin,
            threshold
        });
        const current = target.current;
        observer.observe(current);
        return () => {
            observer.unobserve(current);
        };
    });
};
export default useIntersectionObserver;