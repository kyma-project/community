package tracingsimulation;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;

import java.util.concurrent.ThreadLocalRandom;

/**
 * This sample is based on our official tutorials:
 * <ul>
 *   <li><a href="https://gatling.io/docs/gatling/tutorials/quickstart">Gatling quickstart tutorial</a>
 *   <li><a href="https://gatling.io/docs/gatling/tutorials/advanced">Gatling advanced tutorial</a>
 * </ul>
 */
public class TracingSimulation extends Simulation {


    ChainBuilder traceCall =
        exec(http("Home").get("/"))
            .pause(1);


   

    HttpProtocolBuilder httpProtocol =
        http.baseUrl("https://demo.tracing-poc.berlin.shoot.canary.k8s-hana.ondemand.com/")
            .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
            .acceptLanguageHeader("en-US,en;q=0.5")
            .acceptEncodingHeader("gzip, deflate")
            .userAgentHeader(
                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0"
            );

    ScenarioBuilder trace = scenario("Trace").exec(traceCall);

    {
        setUp(
            trace.injectOpen(rampUsersPerSec(5).to(10).during(6000))
        ).protocols(httpProtocol);
    }
}
