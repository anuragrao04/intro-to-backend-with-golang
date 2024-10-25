from locust import HttpUser, task


class SerialUser(HttpUser):
    @task
    def serialPrimes(self):
        self.client.get("/primes/serial")


# class ParallelUser(HttpUser):
#     @task
#     def parallelPrimes(self):
#         self.client.get("/primes/parallel")
