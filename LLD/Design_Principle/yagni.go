package main

// https://newsletter.ashishps.com/p/8c3c7da7-885b-4a9c-a6e4-70ee02de4772

// What is YAGNI?
// "You Aren't Gonna Need It"
// Always implement things when you actually need them, never when you just foresee that you might need them.
// At its heart, YAGNI is about avoiding over-engineering and focusing strictly on the present requirements.
// The rationale behind YAGNI is simple: every line of code we write comes with a cost.
// It needs to be developed, tested, maintained, and understood by other developers.
// The rationale behind YAGNI is simple: every line of code we write comes with a cost.
// It needs to be developed, tested, maintained, and understood by other developers.

// Over-engineered:
func process_payment1(payment_method string) {
	if payment_method == "credit_card" {
		//complex credit card handling
		return
	} else if payment_method == "paypal" {
		return
		//paypal credit card
	} else if payment_method == "bitcoin" {
		return
		//bitcoin payment handling
	}
}

// YAGNI-aligned:
func process_payment2(payment_method string) {
	if payment_method == "credit_card" {
		//credit card handling
		return
	}
}

// ## Benefits of YAGNI

// 1. Reduced waste
// 2. Simplified codebase
// 3. Faster development
// 4. Improved adaptability

// When YAGNI Might Be Inappropriate
// Like any principle, YAGNI shouldn't be rigidly applied in every situation. There are times when anticipating near-future needs makes sense:

// Well-Known Requirements: If you know with high certainty a feature is coming soon, building some basic support upfront might be wise.

// Performance-Critical Areas: Sometimes, a less-than-optimal but more general solution is necessary initially to ensure performance targets are met.
