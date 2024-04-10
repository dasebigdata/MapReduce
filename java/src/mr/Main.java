package mr;

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Main {

    @SuppressWarnings("unchecked")
    public static void main(String[] args) {
        String inputPath = args[0];
        String outputPath = args[1];
        String loadPath = args[2];
        String mapperClassName = args[3];
        String reducerClassName = args[4];

        try (URLClassLoader loader = new URLClassLoader(new URL[]{new File(loadPath).toURI().toURL()})) {
            Class<?> mapperClass = loader.loadClass(mapperClassName);
            Mapper<Object, Object> mapper = (Mapper<Object, Object>) mapperClass.getConstructor().newInstance();
            Class<?> reducerClass = loader.loadClass(reducerClassName);
            Reducer<Object, Object, Object, Object> reducer = (Reducer<Object, Object, Object, Object>) reducerClass.getConstructor().newInstance();;

            // input
            List<String> lines = new ArrayList<>();
            try (Scanner sc = new Scanner(new FileReader(inputPath))) {
                while (sc.hasNextLine()) {
                    lines.add(sc.nextLine());
                }
            }
            
            // map
            List<Pair<Object, Object>> pairs = new ArrayList<>();
            for (String line : lines) {
                pairs.addAll(mapper.map(line));
            }

            // shuffle
            Map<Object, List<Object>> intermediate = new HashMap<>();
            for (Pair<Object, Object> pair : pairs) {
                List<Object> values = intermediate.get(pair.getKey());
                if (values == null) {
                    intermediate.put(pair.getKey(), new ArrayList<>(){{add(pair.getValue());}});
                } else {
                    values.add(pair.getValue());
                }
            }
            List<Pair<Object, List<Object>>> intermediateList = intermediate.entrySet().stream()
            .map(e -> new Pair<>(e.getKey(), e.getValue()))
            .sorted((e1, e2) -> (e1.getKey().toString()).compareTo(e2.getKey().toString()))
            .collect(Collectors.toList());

            // reduce
            List<Pair<Object, Object>> result = new ArrayList<>();
            for (Pair<Object, List<Object>> pair : intermediateList) {
                result.add(reducer.reduce(pair.getKey(), pair.getValue()));
            }

            // output
            try (BufferedWriter writer = new BufferedWriter(new FileWriter(outputPath))) {
                for (Pair<Object, Object> pair : result) {
                    writer.write(pair.getKey().toString());
                    writer.write("\t");
                    writer.write(pair.getValue().toString());
                    writer.newLine();
                }
            }

        } catch(Exception e) {
            e.printStackTrace();
        }
    }
}
