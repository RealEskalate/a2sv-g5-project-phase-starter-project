
import 'package:application1/features/authentication/domain/entities/user_data.dart';
import 'package:application1/features/authentication/presentation/bloc/auth_bloc.dart';
import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/presentation/bloc/product_bloc.dart';
import 'package:application1/features/product/presentation/pages/home_page.dart';
import 'package:application1/features/product/presentation/widgets/components/product_card.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
class MockProductBloc extends Mock implements ProductBloc {}

class MockAuthBloc extends Mock implements AuthBloc {}

class FakeProductState extends Fake {}

class FakeProductEvent extends Fake  {}

class FakeAuthState extends Fake  {}

class FakeAuthEvent extends Fake  {}

void main() {
  late ProductBloc productBloc;
  late AuthBloc authBloc;

  setUp(() {
    productBloc = MockProductBloc();
    authBloc = MockAuthBloc();

    registerFallbackValue(FakeProductState());
    registerFallbackValue(FakeProductEvent());
    registerFallbackValue(FakeAuthState());
    registerFallbackValue(FakeAuthEvent());
  });
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
  UserEntity tUserEntity =
      const UserEntity(email: 'ley@gmail.com', name: 'ley');
  Widget createWidgetUnderTest() {
    return MaterialApp(
      home: MultiBlocProvider(
        providers: [
          BlocProvider<ProductBloc>.value(value: productBloc),
          BlocProvider<AuthBloc>.value(value: authBloc),
        ],
        child: const Home(),
      ),
    );
  }

  testWidgets('renders Home page with loading state', (tester) async {
    when(() => productBloc.state).thenReturn(ProductLoading());
    when(() => authBloc.state).thenReturn(AuthInitial());

    await tester.pumpWidget(createWidgetUnderTest());

    expect(find.byType(CircularProgressIndicator), findsOneWidget);
  });

  testWidgets('renders Home page with loaded products', (tester) async {

    when(() => productBloc.state).thenReturn(LoadedAllProductState(tProducts));
    when(() => authBloc.state).thenReturn(AuthInitial());

    await tester.pumpWidget(createWidgetUnderTest());

    expect(find.byType(MyCardBox), findsWidgets);
  });

  testWidgets('renders Home page with error state', (tester) async {
    when(() => productBloc.state).thenReturn(const ProductErrorState('Error loading products'));
    when(() => authBloc.state).thenReturn(AuthInitial());

    await tester.pumpWidget(createWidgetUnderTest());

    expect(find.text('Error loading products'), findsOneWidget);
  });

  testWidgets('shows error message when AuthBloc has error', (tester) async {
    when(() => productBloc.state).thenReturn(ProductLoading());
    when(() => authBloc.state).thenReturn(const AuthErrorState(message: 'Authentication error'));

    await tester.pumpWidget(createWidgetUnderTest());
    await tester.pump();

    expect(find.text('Authentication error'), findsOneWidget);
  });

  testWidgets('tapping FAB navigates to product add page', (tester) async {
    when(() => productBloc.state).thenReturn(ProductLoading());
    when(() => authBloc.state).thenReturn(AuthInitial());

    await tester.pumpWidget(createWidgetUnderTest());

    final fab = find.byType(FloatingActionButton);
    expect(fab, findsOneWidget);

    await tester.tap(fab);
    await tester.pumpAndSettle();

    expect(find.byType(Home), findsNothing);
  });

  testWidgets('tapping logout button shows logout dialog', (tester) async {
    when(() => productBloc.state).thenReturn(ProductLoading());
    when(() => authBloc.state).thenReturn(AuthUserLoaded(userEntity: tUserEntity));

    await tester.pumpWidget(createWidgetUnderTest());

    final logoutButton = find.byIcon(Icons.logout);
    expect(logoutButton, findsOneWidget);

    await tester.tap(logoutButton);
    await tester.pumpAndSettle();

    expect(find.text('Logout'), findsOneWidget);
  });

  testWidgets('tapping search icon navigates to search page', (tester) async {
    when(() => productBloc.state).thenReturn(ProductLoading());
    when(() => authBloc.state).thenReturn(AuthInitial());

    await tester.pumpWidget(createWidgetUnderTest());

    final searchIcon = find.byIcon(Icons.search);
    expect(searchIcon, findsOneWidget);

    await tester.tap(searchIcon);
    await tester.pumpAndSettle();

    expect(find.byType(Home), findsNothing);
  });
}

