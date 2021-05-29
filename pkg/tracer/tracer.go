package tracer

//
//func NewJaegerTracer(serviceName, agentHost string) (opentracing.Tracer, io.Closer, error) {
//	config := &config.Configuration{
//		ServiceName: serviceName,
//		Sampler: &config.SamplerConfig{
//			Type:  "const",
//			Param: 1,
//		},
//		Reporter: &config.ReporterConfig{
//			LogSpans:            true,
//			BufferFlushInterval: 1 * time.Second,
//			LocalAgentHostPort:  agentHost,
//		},
//	}
//	tracer, closer, err := config.NewTracer()
//
//	if err != nil {
//		return nil, nil, err
//	}
//	opentracing.SetGlobalTracer(tracer)
//	return tracer, closer, nil
//
//}
