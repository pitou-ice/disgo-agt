FROM baseImage

EXPOSE 8080

CMD [ "./llm-server", \
      "-m", "models/Hermes-2-Pro-Mistral-7B-q4_0.gguf", \
      "-t", "4", \
      "-c", "1024", \
      "-b", "134" ]

