//import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

import 'layout/layout.dart';
import 'layout/example/example.dart';
import 'layout/example/layout_bar.dart';

import 'screens/announcements_screen/announcements_screen.dart';
import 'screens/home_screen/home_screen.dart';
import 'screens/page_not_found_screen/page_not_found.dart';
import 'screens/posts_screen/posts_screen.dart';

void main() {
  runApp(const ProviderScope(child: MyApp()));
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Layout(
        child: MaterialApp(
      title: 'FGGP Boilerplate',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSwatch(primarySwatch: Colors.cyan)
            .copyWith(secondary: Colors.cyanAccent),
      ),
      home: const HomeScreen(), //MyHomePage(title: 'Layout'), //
      routes: {
        PostsScreen.routeName: (ctx) => const PostsScreen(),
        AnnouncementsScreen.routeName: (ctx) => const AnnouncementsScreen(),
      },
      onUnknownRoute: (settings) {
        return MaterialPageRoute<PageNotFoundScreen>(
            builder: (ctx) => const PageNotFoundScreen());
      },
      builder: (context, child) {
        return Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [Expanded(child: child!), LayoutBar()],
        );
      },
    ));
  }
}
