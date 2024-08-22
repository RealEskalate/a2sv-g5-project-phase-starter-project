import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/presentation/bloc/product_bloc.dart';
import 'package:application1/features/product/presentation/pages/product_search_page.dart';
import 'package:application1/features/product/presentation/widgets/components/modal_sheet_widget.dart';
import 'package:application1/features/product/presentation/widgets/components/product_card.dart';
import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

void main() {
  late MockProductBloc mockProductBloc;

  setUp(() {
   
    mockProductBloc = MockProductBloc();
  });

  Widget buildTestableWidget() {
    return MaterialApp(
      home: BlocProvider<ProductBloc>(
        create: (context) => mockProductBloc,
        child: const ProductSearchPage(),
      ),
    );
  }
   List<ProductEntity> tProducts = [
    const ProductEntity(
      id: '1',
      name: 'product1',
      description: 'this is a product',
      price: 23,
      imageUrl:
          'https://i.pinimg.com/564x/ba/86/04/ba86047d55280a343e3c1f0e0868f0e7.jpg',
    ),
    const ProductEntity(
      id: '2',
      name: 'product2',
      description: 'this is a product',
      price: 24,
      imageUrl:
          'https://i.pinimg.com/564x/ba/86/04/ba86047d55280a343e3c1f0e0868f0e7.jpg',
    )
  ];
  testWidgets('shows CircularProgressIndicator when ProductLoading state',
      (WidgetTester tester) async {
    when(() => mockProductBloc.state).thenReturn(ProductLoading());

    await tester.pumpWidget(buildTestableWidget());

    expect(find.byType(CircularProgressIndicator), findsOneWidget);
  });

  testWidgets('shows list of products when LoadedAllProductState',
      (WidgetTester tester) async {
    final products = [
      // Add some mock ProductEntity instances here
    ];

    when(() => mockProductBloc.state).thenReturn(LoadedAllProductState(tProducts));

    await tester.pumpWidget(buildTestableWidget());

    expect(find.byType(MyCardBox), findsNWidgets(products.length));
  });

  testWidgets('shows error message when ProductErrorState',
      (WidgetTester tester) async {
    const errorMessage = 'Something went wrong';
    when(() => mockProductBloc.state).thenReturn(const ProductErrorState(errorMessage));

    await tester.pumpWidget(buildTestableWidget());

    expect(find.text(errorMessage), findsOneWidget);
  });

  testWidgets('shows "No products" text when no products are loaded',
      (WidgetTester tester) async {
    when(() => mockProductBloc.state).thenReturn(ProductInitial());

    await tester.pumpWidget(buildTestableWidget());

    expect(find.text('No products'), findsOneWidget);
  });

  testWidgets('navigates back when back button is pressed',
      (WidgetTester tester) async {
    await tester.pumpWidget(buildTestableWidget());

    await tester.tap(find.byIcon(Icons.arrow_back_ios_rounded));
    await tester.pumpAndSettle();

    // Check navigation pop by using tester state or MockNavigatorObserver if needed.
  });

  testWidgets('opens filter modal when filter button is tapped',
      (WidgetTester tester) async {
    await tester.pumpWidget(buildTestableWidget());

    await tester.tap(find.byIcon(Icons.filter_list_rounded));
    await tester.pumpAndSettle();

    expect(find.byType(ModalSheetComponent), findsOneWidget);
  });
}
