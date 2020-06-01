import React from "react";
import classnames from "classnames";
import Layout from "@theme/Layout";
import Head from "@docusaurus/Head";
import Link from "@docusaurus/Link";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import useBaseUrl from "@docusaurus/useBaseUrl";
import styles from "./styles.module.css";

const features = [
  {
    title: <>Purpose</>,
    imageUrl: "img/conventional-commits.png",
    description: (
      <>
        Tool to make sure your commits are compliant with conventional commits.
      </>
    ),
  },
  {
    title: <>Language agnostic</>,
    imageUrl: "img/undraw_docusaurus_mountain.svg",
    description: (
      <>
        We aim for ease of use across repositories that use different languages.
        This means we ship a Docker image or a binary. For added security you
        can compile from source code.
      </>
    ),
  },
  {
    title: <>Powered by Go</>,
    imageUrl: "img/gopher.svg",
    description: (
      <>
        This project is powered by the Go language. It allows us to
        cross-compile for many different environments and also ship a small
        binary/Docker image that cuts down on bootstrapping and network time.
      </>
    ),
  },
];

function Feature({ imageUrl, title, description }) {
  const imgUrl = useBaseUrl(imageUrl);
  return (
    <div className={classnames("col col--4", styles.feature)}>
      {imgUrl && (
        <div className="text--center">
          <img className={styles.featureImage} src={imgUrl} alt={title} />
        </div>
      )}
      <h3>{title}</h3>
      <p>{description}</p>
    </div>
  );
}

function Home() {
  const context = useDocusaurusContext();
  const { siteConfig = {} } = context;
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="Description will go into a meta tag in <head />"
    >
      <header className={classnames("hero hero--primary", styles.heroBanner)}>
        <div className="container">
          <h1 className="hero__title">{siteConfig.title}</h1>
          <p className="hero__subtitle">{siteConfig.tagline}</p>
          <div className={styles.buttons}>
            <Link
              className={classnames(
                "button button--outline button--secondary button--lg",
                styles.getStarted,
              )}
              to={useBaseUrl("docs/")}
            >
              Get Started
            </Link>
          </div>
        </div>
      </header>
      <main>
        {features && features.length > 0 && (
          <section className={styles.features}>
            <div className="container">
              <div className="row">
                {features.map((props, idx) => (
                  <Feature key={idx} {...props} />
                ))}
              </div>
            </div>
          </section>
        )}
      </main>
    </Layout>
  );
}

export default Home;
