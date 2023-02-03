package jobs

func (r *Runner) registerJobs() {
	SendEmail(r, r.log)
}
