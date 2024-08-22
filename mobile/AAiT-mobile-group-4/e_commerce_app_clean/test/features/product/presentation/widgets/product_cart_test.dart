import 'dart:io';

import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/presentation/widgets/components/product_card.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  setUp(() {
    HttpOverrides.global = null;
  });

  ProductEntity product = const ProductEntity(
    id: '1',
    name: 'product1',
    description: 'this is a product',
    price: 23.0,
    imageUrl:
        'https://i.pinimg.com/564x/ba/86/04/ba86047d55280a343e3c1f0e0868f0e7.jpg',
  );

  testWidgets(
      'MycardBox widget should contain a title, category, price, rating and picture',
      (widgetTester) async {
    //arrange
    await widgetTester.pumpWidget(MaterialApp(
      home: Scaffold(
        body: Column(
          children: [
            Expanded(
              child: MyCardBox(
                product: product,
              ),
            ),
          ],
        ),
      ),
      routes: {
        '/details_page': (context) => const Scaffold(
              body: Text('Successful'),
            )
      },
    ));

    expect(find.byType(Image), findsOneWidget);

    expect(find.text('product1'), findsOneWidget);
    expect(find.text('\$23.0'), findsOneWidget);

    await widgetTester.tap(find.byType(MyCardBox));
    await widgetTester.pumpAndSettle();

    expect(find.text('Successful'), findsOneWidget);
  });
}
