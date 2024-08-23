import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../../auth/presentation/bloc/auth_bloc/auth_bloc.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import '../widgets/product_card.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    context.read<AuthBloc>().add(GetUserEvent());
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is AuthLoggedOut) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('Logged out successfully'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
          Navigator.pushNamedAndRemoveUntil(
            context,
            '/signin',
            (Route<dynamic> route) => false,
          );
        }
      },
      child: Scaffold(
        appBar: AppBar(
          actions: [
            Container(
              padding: const EdgeInsets.all(5.0),
              child: Row(
                children: [
                  ClipRRect(
                    borderRadius: const BorderRadius.all(Radius.circular(20)),
                    child: Image.asset(
                      'assets/img1.jpg',
                      width: 40.0,
                      height: 50.0,
                    ),
                  ),
                  const SizedBox(
                    width: 20.0,
                  ),
                   Column(
                    children: [
                      const Text('July 14,2023',
                          style: TextStyle(fontWeight: FontWeight.w400)),
                      Row(
                        children: [
                          const Text(
                            'Hello, ',
                            style: TextStyle(fontWeight: FontWeight.bold),
                          ),
                          BlocBuilder<AuthBloc, AuthState>(
                            builder: (context, state) {
                              if (state is AuthAuthenticated) {
                                return Text(
                                  state.userDataEntity.name,
                                  style: const TextStyle(fontWeight: FontWeight.bold),
                                );
                              }else{
                                return const Text('User',  style: TextStyle(fontWeight: FontWeight.bold),);
                              }
                             
                            },
                          ),
                        ],
                      )
                    ],
                  )
                ],
              ),
            ),
            const SizedBox(
              width: 130.0,
            ),
            Container(
                decoration: BoxDecoration(
                  border: Border.all(width: 2, color: Colors.grey.shade300),
                  borderRadius: BorderRadius.circular(5),
                ),
                child: InkWell(
                    onTap: () {},
                    splashColor: Colors.grey.shade300,
                    child: Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: Stack(children: [
                        const Icon(Icons.notifications_none_rounded),
                        Positioned(
                            top: 3,
                            right: 5,
                            child: Container(
                              width: 6,
                              height: 6,
                              decoration: BoxDecoration(
                                  borderRadius: BorderRadius.circular(20),
                                  color:
                                      const Color.fromARGB(255, 63, 81, 243)),
                            ))
                      ]),
                    ))),

                      IconButton(
                onPressed: () {
                  context.read<AuthBloc>().add(LogoutEvent());
                },
                icon: const Icon(Icons.logout))
          ],
        ),
        body: Padding(
          padding: const EdgeInsets.all(20),
          child: homeBuilder(context),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            Navigator.pushNamed(context, '/add');
          },
          backgroundColor: Colors.indigoAccent.shade400,
          child: const Icon(
            Icons.add,
            color: Colors.white,
          ),
        ),
      ),
    );
  }
}

Widget homeBuilder(BuildContext context) {
  context.read<ProductBloc>().add(LoadProduct());

  return Column(
    children: [
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          const Text(
            'Available Products',
            style: TextStyle(
                fontSize: 25.0,
                fontWeight: FontWeight.bold,
                fontFamily: AutofillHints.countryName),
          ),
          Container(
              decoration: BoxDecoration(
                border: Border.all(width: 2, color: Colors.grey.shade300),
                borderRadius: BorderRadius.circular(5),
              ),
              child: InkWell(
                  onTap: () {
                    Navigator.pushNamed(context, '/search');
                  },
                  splashColor: Colors.grey.shade300,
                  child: const Padding(
                    padding: EdgeInsets.all(8.0),
                    child: Icon(
                      Icons.search_outlined,
                      color: Colors.black45,
                    ),
                  ))),
        ],
      ),
      const SizedBox(
        height: 10,
      ),
      BlocBuilder<ProductBloc, ProductState>(builder: (context, state) {
        if (state is ProductLoading) {
          return const Center(child: CircularProgressIndicator());
        } else if (state is ProductLoaded) {
          final successState = state;
          return Expanded(
            child: RefreshIndicator(
              onRefresh: () async {
                context.read<ProductBloc>().add(LoadProduct());
              },
              child: GridView.builder(
                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                  crossAxisCount: 1,
                  crossAxisSpacing: 10,
                  mainAxisSpacing: 10,
                  childAspectRatio: 1.3,
                ),
                itemCount: successState.products.length,
                itemBuilder: (BuildContext context, int index) {
                  return ProductCard(
                    product: successState.products[index],
                  );
                },
              ),
            ),
          );
        } else {
          return const Center(child: Text('Error loading products'));
        }
      }),
    ],
  );
}
