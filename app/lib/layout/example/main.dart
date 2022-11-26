import './layout_bar.dart';
import 'package:flutter/material.dart';
import '../layout.dart';

import 'example.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return Layout(
      child: MaterialApp(
          title: 'Flutter Demo',
          debugShowCheckedModeBanner: false,
          theme: ThemeData(
            appBarTheme: const AppBarTheme(
              color: Colors.white,
            ),
            scaffoldBackgroundColor: Colors.grey[200],
            primarySwatch: Colors.blue,
          ),
          builder: (context, child) {
            return Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [Expanded(child: child!), LayoutBar()],
            );
          },
          home: MyHomePage(title: 'Layout')),
    );
  }
}
