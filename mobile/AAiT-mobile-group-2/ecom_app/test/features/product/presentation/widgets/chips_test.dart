
import 'package:ecom_app/features/product/presentation/widgets/chips.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('Chips widget test', (WidgetTester tester) async {
    // Build our app and trigger a frame.
    await tester.pumpWidget(MaterialApp(
      theme: ThemeData(
          // textTheme: GoogleFonts.poppinsTextTheme(),
          primaryColor: const Color.fromARGB(255, 63, 81, 243),
          colorScheme: ColorScheme.fromSeed(
              seedColor: const Color.fromARGB(255, 63, 81, 243))),
      home: const Scaffold(
        body: Chips(
          number: 5,
          selected: true,
        ),
      ),
    ));

    // Verify that our counter starts at 0.
    final chips = tester.widget<Chip>(find.byType(Chip));
    expect(find.byType(Chip), findsOneWidget);
    expect(chips.backgroundColor, const Color.fromARGB(255, 63, 81, 243));
    expect(find.text('5'), findsOneWidget);
  });
}
